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
        <button>Clear</button>
      </div>
    );
  }

  async calcRoute() {
    let data = await this.getRoute();
    //let json = JSON.parse(data);
    console.log(data);
  }

  getRoute() {
    return new Promise((resolve, reject) => {
      request.post({ 
          url: 'http://localhost:8000/calcRoute', 
          form: JSON.stringify(window.points),
        },
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
