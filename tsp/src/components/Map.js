
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

    let numPoints = this.state.points.length;
    let newID = 'P' + (numPoints + 1).toString();

    let newPoints = this.state.points.concat([{id: newID, x: x, y: y}]);
    this.setState({points: newPoints});

    window.points = newPoints;
  }

  render() {

    let points = [];

    for (var i = 0; i < this.state.points.length; i++) {
      points.push(
        <Point key={i} id={this.state.points[i].id} 
          x={this.state.points[i].x} y={this.state.points[i].y} />
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

