import React from 'react'
import {
    Switch,
    Route,
    Redirect
} from "react-router-dom";
import { withGlobalState } from 'react-globally'
import Login from "./pages/login/Login"
import Home from "./pages/home/Home"

const PrivateRoute = ({ component, isAuthed, ...rest }) => {
    return (
        <Route {...rest} exact
            render={(props) => (
                isAuthed ? (
                    <div>
                        {React.createElement(component, props)}
                    </div>
                ) :
                    (
                        <Redirect
                            to={{
                                pathname: '/',
                                state: { from: props.location }
                            }}
                        />
                    )
            )}
        />
    )
}

class App extends React.Component {
    render() {
        return (
            <Switch>
                <Route exact path="/" component={Login} />
                <PrivateRoute exact path="/home" component={Home} isAuthed={this.props.globalState.isLoggedIn} />
            </Switch>
        )
    }
}

export default withGlobalState(App)