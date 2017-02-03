
import React, { Component } from 'react';

class Map extends Component {

  constructor() {
    super(...arguments);
    this.drawPoint = this.drawPoint.bind(this);
  }

  drawPoint() {
    console.log('drawPoint');
  }

  render() {
    return (
      <svg id="map" width="600" height="400" onClick={this.drawPoint}/>
    );
  }
}

export default Map;

