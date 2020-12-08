import React, { Component } from 'react'
import axios from 'axios';
import Urls from './Util/Urls.js';
class PingComponent extends Component {

    constructor() {
        super();
        this.state = {
            pong: 'pending'
        }
    }

    componentWillMount() {
        axios.get(`${Urls.api}/api/config`)
            .then((response) => {
                this.setState(() => {
                    return { pong: response.data.message }
                })
            })
            .catch(function (error) {
                console.log(error);
            });

    }

    render() {
        return <h1>Ping {this.state.pong}</h1>;
    }
}

export default PingComponent; 