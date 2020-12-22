import React, { Component  } from 'react';
import axios from 'axios';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, Label, FormText, FormGroup, Input, Table,UncontrolledAlert} from 'reactstrap';
import Urls from './Util/Urls.js';

class Config extends Component {
	constructor(props) {
		super(props);
		this.state = {
			modal: false,
			config: null,
			success: false,
			error: null,
		}
		this.toggle = this.toggle.bind(this);		
	}

	onSubmitForm = e => {
		e.preventDefault()
		const formData = new FormData(e.target)
		const body = {}
		formData.forEach((value, property) => body[property] = value)
		console.table(body)
		const json = JSON.stringify(body)
		axios.post(`${Urls.api}/config`, json, {
			headers: { 'content-type': 'application/json'}
		})
			.then((res) => {
				this.getConfig().then(data => this.setState({config: data}))
				this.setState({success: true})
			})
			.catch(err => {
				this.setState({error: err})
			});
	}
	
	toggle() {
		this.setState({
		  modal: !this.state.modal
		});
	  }
	componentDidMount() {
		if (!this.state.data) {
			this.getConfig().then(data => this.setState({config: data}))
			.catch(err => {})
		}
	}
	async getConfig() {
		const res = await axios.get(`${Urls.api}/config`);
		
		console.log(res["data"])
		return await res.data;
	}
	
	render() {
		const { config, error, success } = this.state
		const renderErrorAlert = ()=>{
			if(error) {
				return (<div>
						<UncontrolledAlert color="danger">
							I am an alert and I can be dismissed!
						</UncontrolledAlert>
						</div>
					)
				}
		}
		const renderSuccessAlert = ()=>{
			if(success) {
				return (<div>
						<UncontrolledAlert color="success">
							Configuration successfully updated
						</UncontrolledAlert>
						</div>
					)
				}
		}
			return (
				<div>
					{renderErrorAlert()}
					{renderSuccessAlert()}
				<Table bordered >
					<thead>
						
					</thead>
					<tbody>
						<tr>
							<td>Slack URL</td>
							<td>
								<pre>{config ? config.Slack.slackurl : <em>Loading...</em> }</pre>
							</td>
						</tr>
						<tr>
							<td>Slack API</td>
							<td>
								<pre>{config ? config.Slack.slackapi : <em>Loading...</em> }</pre>
							</td>
						</tr>
						<tr>
							<td>Slack Signing Key</td>
							<td>
								<pre>{config ? config.Slack.slacksigning : <em>Loading...</em> }</pre>
							</td>
						</tr>
					</tbody>
				</Table>
				<Button color="danger" onClick={this.toggle}>Edit Config</Button>
      			<Modal isOpen={this.state.modal} toggle={this.toggle} >
				  	<form onSubmit={e => this.onSubmitForm(e)}>
					<ModalHeader toggle={this.toggle}>Modal title</ModalHeader>
					<ModalBody>
						<FormGroup>
							<Label for="slackurl">Slack URL</Label>
							<Input 
								name="slackurl"
								placeholder={config ? config.Slack.slackurl : ""}
								defaultValue={config ? config.Slack.slackurl : ""}
							/>
							<FormText>Enter the URL for the Slack workspace to connect to </FormText>
						</FormGroup>
						<FormGroup>
							<Label for="slackapi">Slack API</Label>
							<Input 
								name="slackapi"
								placeholder={config ? config.Slack.slackapi : ""}
								defaultValue={config ? config.Slack.slackapi : ""}
							/>
							<FormText>Enter the API Key for the Slack workspace  </FormText>
						</FormGroup>
						<FormGroup>
							<Label for="slacksigning">Slack Signing Secret</Label>
							<Input 
								name="slacksigning"
								placeholder={config ? config.Slack.slacksigning : ""}
								defaultValue={config ? config.Slack.slacksigning : ""}
							/>
							<FormText>Enter the <a href="https://api.slack.com/authentication/verifying-requests-from-slack#signing_secrets_admin_page">Signing Secret</a> for the Slack workspace  </FormText>
						</FormGroup>
					</ModalBody>
					<ModalFooter>
						<Button color="primary" type="submit" onClick={this.toggle}>Do Something</Button>{' '}
					</ModalFooter>
					</form>
				</Modal>
				
				{this.state.res && (
						<div className="res-block">
							<h3>Config Saved:</h3>
							<pre>FormData {this.state.res}</pre>
						</div>
					)}
			</div>
			);
		
		}

	
}

export { Config };                 
