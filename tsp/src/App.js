import React, { Component } from 'react';
import Map from './components/Map';
import './App.css';
import request from 'request';

class App extends Component {

  render() {
    return (
      <div className="App">
        <Map />
        <br />
        <button onClick={this.calcRoute.bind(this)}>Calculate shortest route</button>
      </div>
    );
  }

  async calcRoute() {
    let json = await this.getRoute();
    console.log(json);
  }

  getRoute() {
    return new Promise((resolve, reject) => {
      request('http://localhost:8000', 
        function(err, response, body) {
          if (err) {
            reject(err); return;
          }
          resolve(body);
        });
    });
  }
}

export default App;
