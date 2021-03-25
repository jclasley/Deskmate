import React, { Component  } from 'react';
import axios from 'axios';
import { Button } from 'reactstrap';
import Urls from '../Util/Urls.js';

class ZendeskConnect extends Component {
    constructor(props) {
		super(props);
		this.state = {
			connected: false,
			error: null,
		}
	}
    componentDidMount() {
		if (!this.state.connected) {
			this.getConnectedState().then(
                data => {
                    this.setState({connected: data})
                    if (this.state.connected === true) {
                        this.setState({connectedColor: "success"})
                    } else {
                        this.setState({connectedColor: "danger"})
                    }
                })
			.catch(err => {})
		}
    }

	async getConnectedState() {
        const res = await axios.get(`${Urls.api}/zendesk/status`);
        console.log(res.data);

		return await res.data;
	}
    connect = e => {
        e.preventDefault()
        console.log("Connecting to Zendesk...")
        const json = JSON.stringify({"url": window.location.href})
        console.log(json)
		axios.post(`${Urls.api}/zendesk/connect`, json, {
			headers: { 'content-type': 'application/json'}
		})
			.then((res) => {
				this.getConnectedState().then(
                    data => {
                        this.setState({connected: data})
                        if (this.state.connected === true) {
                            this.setState({connectedColor: "success"})
                        } else {
                            this.setState({connectedColor: "danger"})
                        }
                    })
			})
			.catch();
    }
    render(){
        const { connected } = this.state
        return (
            <div>
                <Button color={this.state.connectedColor} onClick={this.connect}>{connected ? <b>Connected: Zendesk</b> : <b>Not Connected: Zendesk</b> }</Button>
            </div>
        )
    }
}
export {ZendeskConnect};