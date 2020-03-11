import React from 'react';
import './Fact.css';

class Fact extends React.Component {

  constructor(props) {
    super(props)

    this.state = {
      text: this.props.text
    }
  }

  render() {
    return (
    <p>{this.state.text}</p>
    );
  }
}

export default Fact;