import React, { Component, PropTypes } from 'react';
import axios from 'axios';
import { Button, Modal, FormGroup, ControlLabel, FormControl, Alert } from 'react-bootstrap/lib';
import Urls from '../util/Urls.js';

class SaveConfigButton extends Component {
  constructor(props) {
    super(props);
    this.state = { showModal: false, slackapi: '', slackurl: '', isLoading: false, errors: [] };
  }

  close() {
    this.setState({ showModal: false });
  }

  open() {
    this.setState({ showModal: true });
  }

  handleChange(key, e) {
    const newState = {};
    newState[key] = e.target.value;
    this.setState(newState);
  }

  checkInput() {
    const errors = [];
    if (this.state.slackapi.length === 0) {
      errors.push('Author cannot be blank.');
    }

    if (this.state.slackurl.length === 0) {
      errors.push('Message cannot be blank.');
    }

    return errors;
  }

  postConfig() {
    const { slackapi,  slackurl} = this.state;
    this.setState({ isLoading: true, errors: [] });
    const errors = this.checkInput();
    if (errors.length === 0) {
      axios.post(`${Urls.api}/api/config`, {
        SlackURL: slackurl,
        SlackAPI: slackapi,
      })
        .then((res) => {
            this.props.updateConfig(res.data);
            this.setState({ isLoading: false, slackapi: '', slackurl: '', showModal: false, errors: [] });
        },
      )
        .catch((err) => {
          this.setState({ isLoading: false, errors: [err.message] });
        },
      );
    } else {
      this.setState({ isLoading: false, errors });
    }
  }

  makeModalErrors() {
    const { errors } = this.state;
    if (errors.length > 0) {
      return (
        <Alert bsStyle="warning">
          {this.state.errors.join('\n')}
        </Alert>
      );
    }

    return <div />;
  }

  render() {
    const { showModal, isLoading } = this.state;
    return (
      <div>
        <Button bsStyle="primary" onClick={this.open.bind(this)}>Create Config</Button>
        <Modal show={showModal} onHide={this.close.bind(this)}>
          <Modal.Header closeButton>
            <Modal.Title>Create Config</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            {this.makeModalErrors()}
            <form>
              <FormGroup>
                <ControlLabel>Slack URL</ControlLabel>
                <FormControl
                  type="text"
                  value={this.state.slackurl}
                  placeholder="Enter Slack URL to connect to"
                  onChange={this.handleChange.bind(this, 'slackurl')}
                />
              </FormGroup>
              <FormGroup>
                <ControlLabel>Slack API</ControlLabel>
                <FormControl
                  type="text"
                  value={this.state.slackapi}
                  placeholder="Enter Slack API key to connect with"
                  onChange={this.handleChange.bind(this, 'slackapi')}
                />
              </FormGroup>
            </form>
          </Modal.Body>
          <Modal.Footer>
            <Button
              onClick={this.createPost.bind(this)}
              disabled={isLoading}
            >
              {isLoading ? 'Submitting...' : 'Submit'}
            </Button>
          </Modal.Footer>
        </Modal>
      </div>
    );
  }
}


export default CreatePostButton;