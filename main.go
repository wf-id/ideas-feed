package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
)

var (
	wg sync.WaitGroup
)

type TemplateData struct {
	Posts []*Post
}

type Post struct {
	Link        string
	Title       string
	Published   time.Time
	Host        string
	Description string
	Taglist     string
}

var (
	feeds = []string{
		"https://www.inoreader.com/stream/user/1005349717/tag/save",
	}

	// Show up to 60 days of posts
	relevantDuration = 90 * 24 * time.Hour

	outputDir  = "docs" // So we can host the site on GitHub Pages
	outputFile = "index.html"

	// Error out if fetching feeds takes longer than a minute
	timeout = time.Minute
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	posts := getAllPosts(ctx, feeds)

	if err := os.MkdirAll(outputDir, 0700); err != nil {
		return err
	}

	f, err := os.Create(path.Join(outputDir, outputFile))
	if err != nil {
		return err
	}
	defer f.Close()

	templateData := &TemplateData{
		Posts: posts,
	}

	if err := executeTemplate(f, templateData); err != nil {
		return err
	}

	return nil
}

// getAllPosts returns all posts from all feeds from the last `relevantDuration`
// time period. Posts are sorted chronologically descending.
func getAllPosts(ctx context.Context, feeds []string) []*Post {
	postChan := make(chan *Post)

	wg.Add(len(feeds))
	for _, feed := range feeds {
		go getPosts(ctx, feed, postChan)
	}

	var posts []*Post
	go func() {
		for post := range postChan {
			posts = append(posts, post)
		}
	}()

	wg.Wait()
	close(postChan)

	// Sort items chronologically descending
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Published.After(posts[j].Published)
	})

	return posts
}

func mycombine(is ...string) {
	for i := 0; i < len(is); i++ {
		fmt.Println(is[i])
	}
}

func getPosts(ctx context.Context, feedURL string, posts chan *Post) {
	defer wg.Done()
	parser := gofeed.NewParser()
	feed, err := parser.ParseURLWithContext(feedURL, ctx)
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range feed.Items {
		published := item.PublishedParsed
		Description := strip.StripTags(item.Description)
		if published == nil {
			published = item.UpdatedParsed
		}
		if published.Before(time.Now().Add(-relevantDuration)) {
			continue
		}
		parsedLink, err := url.Parse(item.Link)
		if err != nil {
			log.Println(err)
		}

		categoryStart := strings.Join(item.Categories[:], " ")

		post := &Post{
			Link:        item.Link,
			Title:       item.Title,
			Published:   *published,
			Host:        parsedLink.Host,
			Description: Description,
			Taglist:     categoryStart,
		}
		posts <- post
	}
}

func executeTemplate(writer io.Writer, templateData *TemplateData) error {
	htmlTemplate := `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>IDEAS | Feed</title>
		<style>
			@import url("https://fonts.googleapis.com/css2?family=Nanum+Myeongjo&display=swap");

			body {
				font-family: "Nanum Myeongjo", serif;
				line-height: 1.7;
				max-width: 600px;
				margin: 50px auto 50px;
				padding: 0 12px 0;
				height: 100%;
			}

			li {
				padding-bottom: 16px;
			}
			.container {
				overflow: hidden;
			  }
			  
			  .filterDiv {
				float: left;
				width: 100%;
				text-align: left;
				margin: 2px;
				display: none; /* Hidden by default */
			  }
			  
			  /* The "show" class is added to the filtered elements */
			  .show {
				display: block;
			  }
			  
			  /* Style the buttons */
			  .btn {
				border: none;
				outline: none;
				padding: 12px 16px;
				background-color: #f1f1f1;
				cursor: pointer;
			  }
			  
			  /* Add a light grey background on mouse-over */
			  .btn:hover {
				background-color: #ddd;
			  }
			  
			  /* Add a dark background to the active button */
			  .btn.active {
				background-color: #666;
				color: white;
			  }
		</style>
		<script>
		filterSelection("all")
function filterSelection(c) {
  var x, i;
  x = document.getElementsByClassName("filterDiv");
  if (c == "all") c = "";
  // Add the "show" class (display:block) to the filtered elements, and remove the "show" class from the elements that are not selected
  for (i = 0; i < x.length; i++) {
    w3RemoveClass(x[i], "show");
    if (x[i].className.indexOf(c) > -1) w3AddClass(x[i], "show");
  }
}

// Show filtered elements
function w3AddClass(element, name) {
  var i, arr1, arr2;
  arr1 = element.className.split(" ");
  arr2 = name.split(" ");
  for (i = 0; i < arr2.length; i++) {
    if (arr1.indexOf(arr2[i]) == -1) {
      element.className += " " + arr2[i];
    }
  }
}

// Hide elements that are not selected
function w3RemoveClass(element, name) {
  var i, arr1, arr2;
  arr1 = element.className.split(" ");
  arr2 = name.split(" ");
  for (i = 0; i < arr2.length; i++) {
    while (arr1.indexOf(arr2[i]) > -1) {
      arr1.splice(arr1.indexOf(arr2[i]), 1);
    }
  }
  element.className = arr1.join(" ");
}

// Add active class to the current control button (highlight it)
var btnContainer = document.getElementById("myBtnContainer");
var btns = btnContainer.getElementsByClassName("btn");
for (var i = 0; i < btns.length; i++) {
  btns[i].addEventListener("click", function() {
    var current = document.getElementsByClassName("active");
    current[0].className = current[0].className.replace(" active", "");
    this.className += " active";
  });
}

		</script>
	</head>
	<body>
	<P></P><cite><a href="https://wakeforestid.com/">home</a></cite>
		<h1>Latest Literature</h1>
		<div id="myBtnContainer">
		<button class="btn active" onclick="filterSelection('all')"> Show all</button>
		<button class="btn" onclick="filterSelection('amr')"> AMR</button>
		<button class="btn" onclick="filterSelection('fungal')"> Fungal</button>
		<button class="btn" onclick="filterSelection('std')"> STD</button>
		<button class="btn" onclick="filterSelection('surveillance')"> Surveillance</button>
		</div>

		<div class="container">
		<ol>
			{{ range .Posts }}<div class="filterDiv {{ .Taglist}}">
			  <li><a href="{{ .Link }}">{{ .Title }}</a><p>{{ .Description }}</p> ({{ .Host }})</li>
			  </div>
			{{ end }}
		</ol>
		</div>

		

		<footer>
		    <p><a href="https://feed.wakeforestid.com/bibliography.bibtex" download>Download this bibliography</a></p>
		    <p><a href="https://github.com/wf-id/ideas-feed">What is this?</a></p>
			<p><a href="https://www.zotero.org/groups/4900647/ideas-feed/library">Zotero Library</a></p>
			<p><a href="https://www.researchrabbit.ai/">Research Rabbit</a></p>
			<p><a href="https://github.com/jamesroutley/news.routley.io">What is this based on</a></p>
			<p><a href="https://wakeforestid.com">Main Website</a></p>
		</footer>
	</body>
</html>
`

	tmpl, err := template.New("webpage").Parse(htmlTemplate)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(writer, templateData); err != nil {
		return err
	}

	return nil
}
