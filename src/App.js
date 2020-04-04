import React, { Component } from "react";
import Autocomplete from "./Autocomplete";
import './App.css';

export default class App extends Component {

  componentDidMount() {
    fetch("/hub")
        .then(res => res.json())
        .then(result =>  this.setState({ revs: result}))
}
  
  
render(){
  return (
    <div className="App">
      {this.state && this.state.revs.length > 0 &&
        <Autocomplete
        suggestions={this.state.revs}
        />
      }
    </div>
  );
}
}