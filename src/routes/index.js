import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import HelloWorld from "../components/HelloWord";
import Company from "../components/Company";
import Position from "../components/Position";
import Candidate from "../components/Candidate";
import Folder from "../components/Folder";
import Task from "../components/Task";


import CMenu from "../components/Menu";
/*
export default class CRouter extends React.Component {
  render() {
    return (
      <Router>
        <Route path="/" components={HelloWorld} />
      </Router>
    );
  }
}
*/
const CRouter = () => (
  <Router>
    <div>
      <Route path="/" component={CMenu} />
      <Route path="/company" component={Company} />
      <Route path="/position" component={Position} />
      <Route path="/candidate" component={Candidate} />
      <Route path="/task" component={Task} />
      <Route path="/folder" component={Folder} />
    </div>
  </Router>
);

export default CRouter;
