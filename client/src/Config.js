import React, { Component } from 'react';
import axios from 'axios';
import { Button, Form, Label, FormText, FormGroup, Input} from 'reactstrap';
import Urls from './Util/Urls.js';


class ConfigEditor extends Component {
	constructor(props) {
		super(props);
		this.state = {
			config: {},
		}
		this.submitForm = this.submitForm.bind(this);
		this.handleChange = this.handleChange.bind(this);
	}

	async componentDidMount() {
		const response = await axios.get(`${Urls.api}/config`);
		this.setState( { config: response.data })
		console.log(response)
	}

	handleChange(event) {
		const target = event.target;
		const value = target.type === 'checkbox' ? target.checked : target.value;
		const name = target.name;

		this.setState({
			[name]: value
		});
	}
	submitForm(e) {
		e.preventDefault();
		const data = new FormData(e.target);
		
		if (this.state.slackapi === ""){
			this.config.Slack.slackapi = this.state.slackapi;
		}
		if (this.state.slackurl === ""){
			this.config.Slack.slackurl = this.state.slackurl;
		}
		axios({
			method: 'post',
			url: `${Urls.api}/config`,
			data: data ,
			headers: { 'content-type': 'application/json'}
		})
			.then((res) => {
				this.setState({ res: stringifyFormData(data) });
			},
		)
			.catch((err) => {
				
			},
		);
	}
	render() {
		const {config} = this.state
		if (config.Slack) {
			return (
			<div>
				<Form onSubmit={this.submitForm}>
					<FormGroup>
						<Label for="slackurl">Slack URL</Label>
						<Input 
							name="slackurl"
							placeholder={config.Slack.slackurl}
							defaultValue={config.Slack.slackurl}
							onChange={this.handleChange}/>
						<FormText>Enter the URL for the Slack workspace to connect to {this.slackurl} </FormText>
					</FormGroup>
					<FormGroup>
						<Label for="slackapi">Slack API</Label>
						<Input 
							name="slackapi"
							placeholder={config.Slack.slackapi}
							defaultValue={config.Slack.slackapi}
							onChange={this.handleChange}/>
						<FormText>Enter the API Key for the Slack workspace {this.slackapi} </FormText>
					</FormGroup>
					<Button type="submit">Save & Close</Button>
		
				</Form>
				{this.state.res && (
						<div className="res-block">
							<h3>Config Saved:</h3>
							<pre>FormData {this.state.res}</pre>
						</div>
					)}
			</div>
			);
		}
		return <div>Loading config...</div>
	}
}


function stringifyFormData(fd) {
	const data = {};
	for (let key of fd.keys()) {
		data[key] = fd.get(key);
	}
	return JSON.stringify(data, null, 2);
}
export { ConfigEditor };                 