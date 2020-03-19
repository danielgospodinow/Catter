import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-globally'
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';

const initialState = {
    isLoggedIn: false
}

ReactDOM.render(
    <BrowserRouter>
        <Provider globalState={initialState}>
            <App />
        </Provider>
    </BrowserRouter>,
    document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
