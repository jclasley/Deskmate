import React, { Component  } from 'react';
import axios from 'axios';
import { Button } from 'reactstrap';
import Urls from '../Util/Urls.js';

class SlackConnect extends Component {
    constructor(props) {
		super(props);
		this.state = {
			connected: false,
            connectedColor: "danger", // should have an initial state for each thing
			error: null,
		}
	}
    async componentDidMount() {
        if (!this.state.connected) {
            try {
                const { data } = await axios.get(`${Urls.api}/slack/status`);
                this.setState({ connected: data });
                this.setState({ connectedColor: data ? "success" : "danger" });
            } catch (err) {
                console.trace(err);
            }
		}
    }

	async getConnectedState() {
        const res = await axios.get(`${Urls.api}/slack/status`);
        console.log(res);

        // because this is 'await'ed already, there is no need to return it as a promise.
        // async/await avoids having to do the promise chain, we can just await this function when it's called in the future
        return res.data;
	}
    connect = e => {
        e.preventDefault()
        console.log("Connecting to Slack...")
		axios.get(`${Urls.api}/slack/connect`)
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
                <Button
                    color={this.state.connectedColor}
                    onClick={this.connect}
                    data-testid="slack_connect"
                >
                    {connected ? <b>Connected: Slack</b> : <b>Not Connected: Slack</b>}
                </Button>
            </div>
        )
    }
}
export {SlackConnect};