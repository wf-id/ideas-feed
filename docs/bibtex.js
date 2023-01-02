const { jsonToBibtex } =  require("@devisle/reference-js");

const content = {
    references: [
      {
        type: "ARTICLE",
        key: "smit54",
        AUTHOR: "J. G. Smith and H. K. Weston",
        TITLE: "Nothing Particular in this Years History",
        YEAR: "1954",
        JOURNAL: "J. Geophys. Res.",
        VOLUME: "2",
        PAGES: "14-15"
      },
      {
        type: "BOOK",
        key: "colu92",
        AUTHOR: "Christopher Columbus",
        TITLE: "How I Discovered America",
        YEAR: "1492",
        PUBLISHER: "Hispanic Press",
        ADDRESS: "Barcelona"
      }
    ]
  };

  function useEffect(content) {
    return jsonToBibtex(JSON.stringify(content), "references");
    }

    console.log(useEffect(content));

