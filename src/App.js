import React, { useEffect, useState } from "react";
import { jsonToBibtex } from "@devisle/reference-js";
import "./styles.css";

const content = {
  references: [{"key":"KA5B7HA9","version":0,"itemType":"journalArticle","creators":[{"firstName":"P.","lastName":"Grassberger","creatorType":"author"}],"tags":[],"issue":"2","pages":"157-172","date":"1983-04-01","ISSN":"0025-5564","DOI":"10.1016/0025-5564(82)90036-0","url":"https://www.sciencedirect.com/science/article/pii/0025556482900360","abstractNote":"Scaling laws are formulated for the behavior of a space-dependent fluctuating general epidemic process near the critical point. Restricted to stationary properties, these laws describe also the critical behavior of random percolation. Monte Carlo calculations are used to estimate the critical exponents and the universal shape of the propagating wave, in the case of 2-dimensional space.","title":"On the critical behavior of the general epidemic process and dynamical percolation","journalAbbreviation":"Mathematical Biosciences","volume":"63","publicationTitle":"Mathematical Biosciences","language":"en","libraryCatalog":"ScienceDirect","accessDate":"2023-01-01T14:41:53Z"}]
};

export default function App() {
  const [example, setExample] = useState("");

  useEffect(() => {
    jsonToBibtex(JSON.stringify(content), "references")
      .then(data => setExample(data))
      .catch(error => setExample(error.message));
  }, []);

  return (
    <div className="App">
      <pre>
        <code>{example}</code>
      </pre>
    </div>
  );
}
