import React from 'react'
import FacebookLogin from 'react-facebook-login'
import './Login.css'
import { Redirect } from 'react-router-dom'
import { withGlobalState } from 'react-globally'

class Login extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            loggedIn: false
        }
    }

    responseFacebook(response) {
        if (response && new Date().getTime() / 1000 < response.data_access_expiration_time) {
            this.props.setGlobalState({
                isLoggedIn: true
            })
            this.setState({
                loggedIn: true
            })
        }
    }

    render() {
        return (
            !this.state.loggedIn
                ? (
                    <div className="App-header">
                        <FacebookLogin
                            appId="149050146339424"
                            autoLoad={true}
                            fields="name,email,picture"
                            callback={this.responseFacebook.bind(this)} />
                    </div>)
                //TODO: Pass parameters to Home somehow (username, etc).. or use global state
                : (<Redirect to="/home" />)
        )
    }
}

export default withGlobalState(Login)
