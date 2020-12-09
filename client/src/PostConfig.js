import React, { Component } from 'react';
import axios from 'axios';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, Form, Label, FormText, FormGroup, Input} from 'reactstrap';
import Urls from './Util/Urls.js';


class CreateConfigButton extends Component {
  constructor(props) {
    super(props);
      this.state = {
        slackapi:"",
        slackurl:"",
      modal: false,
    }
    this.toggle = this.toggle.bind(this);
    this.submitForm = this.submitForm.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }
  toggle() {
    this.setState({
      modal: !this.state.modal
    });
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
    console.log(this.state.slackapi)
    this.toggle();
    axios.post(`${Urls.api}/config`, {
      slackapi: this.state.slackapi,
      slackurl: this.state.slackurl,
    }, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
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
    return (
    <div>
      <Button color="danger" onClick={this.toggle}>Create Config</Button>
      <Modal isOpen={this.state.modal} toggle={this.toggle}>
      <Form onSubmit={this.submitForm}>
        <ModalHeader toggle={this.toggle}>Create Configuration</ModalHeader>
        <ModalBody>
            <FormGroup>
              <Label for="slackurl">Slack URL</Label>
              <Input 
                name="slackurl"
                onChange={this.handleChange}/>
              <FormText>Enter the URL for the Slack workspace to connect to</FormText>
            </FormGroup>
            <FormGroup>
              <Label for="slackapi">Slack API</Label>
              <Input 
                name="slackapi"
                onChange={this.handleChange}/>
              <FormText>Enter the API Key for the Slack workspace</FormText>
            </FormGroup>
        </ModalBody>
        <ModalFooter>
        <Button type="submit">Save & Close</Button>
       
        </ModalFooter>
        </Form>
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

function stringifyFormData(fd) {
  const data = {};
	for (let key of fd.keys()) {
  	data[key] = fd.get(key);
  }
  return JSON.stringify(data, null, 2);
}
export default CreateConfigButton;                 