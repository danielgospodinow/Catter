import React from 'react';
import logo from '../../images/cat2.png';
import Fact from "../../components/fact/Fact"
import './Home.css';

class Home extends React.Component {

  constructor(props) {
    super(props)

    this.state = {
      currentFact: "Loading ..."
    }
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} width="12%" className="App-logo" alt="logo" />

          <div className="facts-container">
            <Fact text={this.state.currentFact}></Fact>
          </div>
        </header>
      </div>
    )
  }
}

export default Home;
