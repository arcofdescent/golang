
import React, { Component } from 'react';
import Point from './Point';

class Map extends Component {

  constructor() {
    super(...arguments);
    this.drawPoint = this.drawPoint.bind(this);

    this.state = {
      points: [],
    };
  }

  drawPoint(event) {
    console.log('drawPoint');
    let rect = event.target.getBoundingClientRect();
    let x = event.clientX - rect.left;
    let y = event.clientY - rect.top;

    let newPoints = this.state.points.concat([{x: x, y: y}]);
    this.setState({points: newPoints});
  }

  render() {

    let points = [];

    for (var i = 0; i < this.state.points.length; i++) {
      points.push(
        <Point key={i} x={this.state.points[i].x} y={this.state.points[i].y} />
      );
    }


    return (
      <svg id="map" width="600" height="400" onClick={this.drawPoint}>
        {points}
      </svg>
    );
  }
}

export default Map;

