import React, { Component } from 'react';

class Point extends Component {

  constructor() {
    super(...arguments);
    this.radius = 10;
  }

  render() {

    let style = {
      stroke: '#000',
    };

    return (
      <circle cx={this.props.x} cy={this.props.y} r={this.radius} style={style} />
    );
  }

}

export default Point;

