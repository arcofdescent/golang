import React, { Component } from 'react';
import Map from './components/Map';
import Result from './components/Result';
import './App.css';
import request from 'request';

class App extends Component {

  constructor() {
    super(...arguments);

    this.state = {
      points: [],
      result: {
        Route: [],
      },
    };

    this.addPoint = this.addPoint.bind(this);
  }

  render() {
    return (
      <div className="App">
        <Map points={this.state.points} addPoint={this.addPoint} result={this.state.result} />
        <br />
        <button onClick={this.calcRoute.bind(this)}>Calculate shortest route</button>
        <button>Clear</button>
        <Result res={this.state.result} />
      </div>
    );
  }

  addPoint(p) {
    console.log('adding point');

    let newPoints = this.state.points.concat([p]);
    this.setState({points: newPoints});
  }

  async calcRoute() {
    let data = await this.getRoute();
    let json = JSON.parse(data);
    console.log({json});
    this.setState({result: json});
  }

  getRoute() {
    return new Promise((resolve, reject) => {
      request.post({ 
          url: 'http://localhost:8000/calcRoute', 
          form: JSON.stringify(this.state.points),
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
