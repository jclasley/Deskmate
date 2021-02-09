import React, { Component  } from 'react';
import Urls from '../Util/Urls.js';
import axios from 'axios';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, Label, FormText, FormGroup, Input, Table, Alert  } from 'reactstrap';


class Tags extends Component {

	intervalID;
    constructor(props) {
		super(props);
		this.state = {
			modal: false,
        success: false,
        error: null,
        alertColor: "",
        alertVisible: false,
        alertMessage: "",
        data: "",
		}
		this.toggle = this.toggle.bind(this);		
	}

	componentDidMount() {
	  this.getData();
	}

	componentWillUnmount() {
	  clearTimeout(this.intervalID);
    }
    
    onSubmitForm = e => {
		e.preventDefault()
		const formData = new FormData(e.target)
		const body = {}
		formData.forEach((value, property) => body[property] = value)
		console.table(body)
        const json = JSON.stringify(body)
        
		axios.post(`${Urls.api}/tags`, json, {
			headers: { 'content-type': 'application/json'}
		})
			.then((res) => {
				this.getData().then(data => this.setState({data: data}))
				this.setState({alertVisible: true, alertMessage: "Tag Created Successfully", alertColor: "success"}, ()=> {window.setTimeout(()=>{this.setState({alertVisible:false})},8000)});
			})
			.catch(error => {
                console.log(error)
                const err = JSON.stringify(error)
				this.setState({alertVisible: true, alertMessage: err, alertColor: "danger"}, ()=> {window.setTimeout(()=>{this.setState({alertVisible:false})},8000)});
            });
            
        
	}
	

	getData = () => {
	  fetch(`${Urls.api}/tags`)
		.then(response => response.json())
		.then(data => {
          this.setState({ data: data});
          console.log(data);
		  // call getData() again in 5 seconds
		  this.intervalID = setTimeout(this.getData.bind(this), 30000);
		});
    }
    
    

    getChannels = () => {
        fetch(`${Urls.api}/slack/channels`)
          .then(response => response.json())
          .then(data => {
            this.setState({ slackusers: data});
          });
    }

	handleClick = tagID => {
		const requestOptions = {
		  method: 'DELETE'
		};
	  
		fetch(`${Urls.api}/tags/`  +  tagID, requestOptions).then((response) => {
		  return response.json();
		}).then((result) => {
		  this.getData()
		});
	}
    toggle() {
		this.setState({
		  modal: !this.state.modal
		});
	  }
	render(){
	
		
		return (

			<div>
                <Alert color={this.state.alertColor} isOpen={this.state.alertVisible} toggle={(e) => this.setState({alertVisible: false})}> {this.state.alertMessage} </Alert>
			<Table>
				<thead>
				<tr>
				<th>Tag</th>
				<th>Slack ID</th>
				<th>Group ID</th>
				<th>Channel</th>
                <th>NotificationType</th>
                <th>Added</th>
				</tr>
			</thead>
			<tbody>
			{this.state.data &&
			 this.state.data.map((item, index) => (
				<tr>
				<th scope="row">{item.Tag}</th>
				<td>{item.SlackID}</td>
				<td>{item.GroupID}</td>
                <td>{item.Channel}</td>
                <td>{item.NotificationType}</td>
                <td>{item.Added}</td>

				<td>
				<button onClick={() => { this.handleClick(item.ID) }} className="delete-btn">Delete</button>
				</td>
			</tr>
			))}
			
			</tbody>
		  </Table>

          <Button color="primary" onClick={this.toggle}>Add Tag</Button>
      			<Modal isOpen={this.state.modal} toggle={this.toggle} >
				  	<form onSubmit={e => this.onSubmitForm(e)}>
					<ModalHeader toggle={this.toggle}>Add Tag</ModalHeader>
					<ModalBody>
						<FormGroup>
							<Label for="tag">Tag</Label>
							<Input 
								name="tag"
							/>
							<FormText>Enter the tag notifications will be alerted on </FormText>
						</FormGroup>
						<FormGroup>
							<Label for="users">Users</Label>
							<UserDropdown />
							<FormText>Select your username from this list</FormText>
						</FormGroup>
						<FormGroup>
							<Label for="group">Slack Group</Label>
							<GroupDropdown />
							<FormText>Select the Group that this tag applies to  </FormText>
						</FormGroup>
                        <FormGroup>
							<Label for="channel">Slack Channel</Label>
							<ChannelDropdown />
							<FormText>Select the Channel that this tag applies to  </FormText>
						</FormGroup>
                        <FormGroup>
							<Label for="notifyType">Notification Type</Label>
							<Input type="select" name="notificationType">
                                <option value="new">New Tickets</option>
                                <option value="breaches">SLA Breaches</option>
                                <option value="updates">Ticket Updates</option>
                            </Input>
							<FormText>Select the Channel that this tag applies to  </FormText>
						</FormGroup>
					</ModalBody>
					<ModalFooter>
						<Button color="primary" type="submit" onClick={this.toggle}>Add Tag</Button>{' '}
					</ModalFooter>
					</form>
				</Modal>
			</div>
		)
	}
}

function UserDropdown(){

    
    const [loading, setLoading] = React.useState(true);
    const [items, setItems] = React.useState([]);
    const [value, setValue] = React.useState("");
    React.useEffect(() => {
        let unmounted = false;
        async function getUsers() {
          const response = await fetch(`${Urls.api}/slack/users`);
          const body = await response.json();
          if (!unmounted) {

            if (body != null){
                setItems(body.map(({ UserName, ID }) => ({ label: UserName, value: ID })));
            }
            
          setLoading(false);
          }
        }
        getUsers();
        return () => {
            unmounted = true;
          };
      }, []);
      return (
        <Input type="select"  disabled={loading} value={value}
        onChange={e => setValue(e.currentTarget.value)} name="slackID">
        {items.map(({ label, value }) => (
            <option key={value} value={value}>
            {label}
            </option>
        ))}
        </Input>
      );
}
function GroupDropdown(){

    
    const [loading, setLoading] = React.useState(true);
    const [items, setItems] = React.useState([]);
    const [value, setValue] = React.useState("");
    React.useEffect(() => {
        let unmounted = false;
        async function getGroups() {
          const response = await fetch(`${Urls.api}/slack/groups`);
          const body = await response.json();
          if (!unmounted) {
              if (body != null){
                setItems(body.map(({ GroupName, ID }) => ({ label: GroupName, value: ID })));
              
                setLoading(false);
              }
          }
        }
        getGroups();
        return () => {
            unmounted = true;
          };
      }, []);
      return (
        <Input type="select"  disabled={loading} value={value}
        onChange={e => setValue(e.currentTarget.value)} name="groupID">
        {items.map(({ label, value }) => (
            <option key={value} value={value}>
            {label}
            </option>
        ))}
        </Input>
      );
}

function ChannelDropdown(){

    
    const [loading, setLoading] = React.useState(true);
    const [items, setItems] = React.useState([]);
    const [value, setValue] = React.useState("");
    React.useEffect(() => {
        let unmounted = false;
        async function getChannels() {
          const response = await fetch(`${Urls.api}/slack/channels`);
          const body = await response.json();
          if (!unmounted) {
            if (body != null){
              body.sort(function(a, b) {
                var textA = a.ChannelName.toUpperCase();
                var textB = b.ChannelName.toUpperCase();
                return (textA < textB) ? -1 : (textA > textB) ? 1 : 0;
              });
              setItems(body.map(({ ChannelName, ID }) => ({ label: ChannelName, value: ID })));
            }
          setLoading(false);
          }
        }
        getChannels();
        return () => {
            unmounted = true;
          };
      }, []);
      return (
        <Input type="select" disabled={loading} value={value}
        onChange={e => setValue(e.currentTarget.value)} name="channel">
        {items.map(({ label, value }) => (
            <option key={value} value={value}>
            {label}
            </option>
        ))}
        </Input>
      );
}

export {Tags };