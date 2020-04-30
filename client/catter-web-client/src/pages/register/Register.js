import React from 'react'
import './Register.css'
import { Button, TextField, Snackbar } from '@material-ui/core'
import { Route } from 'react-router-dom'
import { withGlobalState } from 'react-globally'
import PasswordValidator from 'password-validator'
import * as bcrypt from 'bcryptjs'

const errorMessages = {
    min: "too short",
    max: "too long",
    uppercase: "missing uppercase characters",
    lowercase: "missing lowercase characters",
    digits: "missing digits",
    spaces: "containing spaces"
}

const initialState = {
    error: {
        failedLogin: false,
        message: ""
    },

    user: {
        firstName: "",
        lastName: "",
        email: "",
        password: "",
        confirmPassword: ""
    }
}

class Register extends React.Component {
    constructor(props) {
        super(props)

        this.state = initialState
        this.handleChange = this.handleChange.bind(this)

        this.validator = new PasswordValidator()
        this.validator.is().min(8)
            .is().max(30)
            .has().uppercase()
            .has().lowercase()
            .has().digits()
            .has().not().spaces()
    }

    onRegisterClicked(history) {
        const invalidPasswordIdentifiers = this.validator.validate(this.state.user.password, { list: true })
        if (invalidPasswordIdentifiers.length === 0) {
            const that = this

            bcrypt.hash(this.state.user.password, 10, function (err, hashedPassword) {
                if (!err) {
                    fetch(process.env.REACT_APP_ACCOUNT_SERVICE_URL + "/account", {
                        method: 'POST',
                        headers: { 'Content-Type': 'application/json' },
                        body: {
                            "Name": that.state.user.firstName + that.state.user.lastName,
                            "Email": that.state.user.email,
                            "Password": hashedPassword
                        }
                    })
                        .then(res => res.json())
                        .then(res => {
                            console.log("Created account: " + res.json())
                            history.push("/")
                        })
                        .catch(err => {
                            console.log("Failed to post account to backend! Error: " + err)
                            that.setError("Failed to create account, there's a problem on the server!", true)
                        })
                } else {
                    console.log("Error while hashing function! " + err.message)
                    that.setError("Something fishy happened, try again!", true)
                }
            })
        } else {
            const passwordErrorMessage = this.constructPasswordErrorMessage(invalidPasswordIdentifiers)
            console.log(passwordErrorMessage)
            this.setError(passwordErrorMessage, true)
        }
    }

    render() {
        return (
            <div className="Container">
                <div className="TextFields">
                    <TextField fullWidth={true} value={this.state.user.firstName} onChange={this.handleChange} id="firstName" label="First Name" variant="outlined" />
                    <TextField fullWidth={true} value={this.state.user.lastName} onChange={this.handleChange} id="lastName" style={{ marginTop: "15px" }} label="Last Name" variant="outlined" />
                    <TextField fullWidth={true} value={this.state.user.email} onChange={this.handleChange} id="email" style={{ marginTop: "15px" }} label="Email" variant="outlined" />
                    <TextField fullWidth={true} value={this.state.user.password} onChange={this.handleChange} id="password" style={{ marginTop: "15px" }} label="Password" type="password" variant="outlined" />
                    <TextField fullWidth={true} value={this.state.user.confirmPassword} onChange={this.handleChange} id="confirmPassword" style={{ marginTop: "15px" }} label="Confirm Password" type="password" variant="outlined" />
                </div>

                <div className="Buttons" style={{ marginTop: "3%" }}>
                    <Route render={({ history }) => (
                        <Button fullWidth={true} variant="contained" color="secondary" onClick={() => this.onRegisterClicked(history)}>Create Account</Button>
                    )} />
                </div>

                <Snackbar
                    anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
                    open={this.state.error.failedLogin}
                    message={this.state.error.message}
                    autoHideDuration={5000}
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

    constructPasswordErrorMessage(invalidPasswordIdentifiers) {
        return "Whoops, the password is " + invalidPasswordIdentifiers.map(indentifier => errorMessages[indentifier]).join(", ")
    }

    handleChange(event) {
        event.persist()
        this.setState((state) => state.user[event.target.id] = event.target.value)
    }
}

export default withGlobalState(Register)