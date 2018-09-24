import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import { Provider } from "react-redux"
import { createStore } from "redux"
import reducer from "./reducer"
const root = document.getElementById('root')

const defaultState = {
    samples: [
        {
            id: 1,
            encodedImage: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAAv0lEQVR4nM2RIRPCMAyFo4ftj2A2RWPRwzI9yzwSj+00SJieHXqWzmJXPWRT2HW4ZlMIkrvmrt81eekD+N+IU1kULbU6ZmBDdsyGgZWjIXXTVwwU9EDECJaCG1qoGUX3OejUQspoAlJdWrpxiwCsvVpyOQcFWbNH1B371pntcG577hMg3/ly7pNp0ZmbWSmjLrgrx6YgjTMBJPTlYqzBAL51KldpO2x6Ckcdv27SlXPF+/l6HliVYqNqlbBm/jY+RNpgFnTkHPkAAAAASUVORK5CYII=",
            label: 5
        }
    ]
}

const store = createStore(reducer, defaultState)

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>
    , root);
registerServiceWorker();
