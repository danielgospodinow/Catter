import React from 'react'
import './Login.css'
import logo from '../../images/cat2.png'
import { Route } from 'react-router-dom'
import { withGlobalState } from 'react-globally'
import { Button, TextField, Snackbar } from '@material-ui/core'
import * as bcrypt from 'bcryptjs'

const initialState = {
    error: {
        failedLogin: false,
        message: ""
    },

    user: {
        email: "",
        password: ""
    }
}

class Login extends React.Component {
    constructor(props) {
        super(props)

        this.state = initialState
        this.handleChange = this.handleChange.bind(this)
    }

    onLoginClicked(history) {
        fetch(process.env.REACT_APP_ACCOUNT_SERVICE_URL + "/account/" + this.state.user.email)
            .then(res => res.json())
            .then(account => {
                bcrypt.compare(this.state.user.password, account.Password, (err, equalPasswords) => {
                    if (!err) {
                        if (equalPasswords) {
                            console.log("User with email address '" + account.Email + "' logged in succesfully!")

                            this.props.setGlobalState({
                                isLoggedIn: true
                            })
                            history.push("/home")
                        } else {
                            console.log("Password missmatch!")
                            this.setError("Incorrect email or password!", true)
                        }
                    } else {
                        console.log("Error while checking hashed password! " + err)
                        this.setError("Something fishy happened, try again!", true)
                    }
                })
            })
            .catch(err => {
                console.log(err)
                this.setError("Failed to retrieve account!", true)
            })
    }

    onRegisterClicked(history) {
        history.push("/register")
    }

    render() {
        return (
            <div className="Container">
                <div>
                    <img src={logo} height="200px"></img>
                </div>

                <div className="TextFields" style={{ marginTop: "3%" }}>
                    <TextField fullWidth={true} value={this.state.user.email} onChange={this.handleChange} id="email" label="Email" variant="outlined" />
                    <TextField fullWidth={true} value={this.state.user.password} onChange={this.handleChange} style={{ marginTop: "15px" }} id="password" label="Password" type="password" variant="outlined" />
                </div>

                <div className="Buttons" style={{ marginTop: "3%" }}>
                    <Route render={({ history }) => (
                        <Button fullWidth={true} variant="contained" color="primary" onClick={() => this.onLoginClicked(history)}>Login</Button>
                    )} />

                    <Route render={({ history }) => (
                        <Button fullWidth={true} style={{ marginTop: "2%" }} variant="contained" color="secondary" onClick={() => this.onRegisterClicked(history)}>Register</Button>
                    )} />
                </div>

                <Snackbar
                    anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
                    open={this.state.error.failedLogin}
                    message={this.state.error.message}
                    autoHideDuration={2000}
                    onClose={() => this.setError("", false)}
                />
            </div>
        )
    }

    setError(message, showError) {
        this.setState({
            error: {
                failedLogin: showError,
                message: message
            }
        })
    }

    handleChange(event) {
        event.persist()
        this.setState((state) => state.user[event.target.id] = event.target.value)
    }
}

export default withGlobalState(Login)
