import React from "react";
import { BrowserRouter as Router, Route, NavLink } from "react-router-dom";
import List from "./List";
import Customer from "./Customer";

function App() {
  return (
    <body>
      <Router>
        <header>
          <nav>
            <ul>
              <li>
                <NavLink to="/">Home</NavLink>
                <NavLink to="/">Customers</NavLink>
              </li>
            </ul>
          </nav>
        </header>
        <main style={{ padding: "2rem" }}>
          <Route exact path="/" component={List} />
          <Route path="/customer/:id" component={Customer} />
        </main>
      </Router>
    </body>
  );
}

export default App;
