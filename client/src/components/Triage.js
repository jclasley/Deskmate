import React, { Component  } from 'react';
import Urls from '../Util/Urls.js';
import { Table } from 'reactstrap';


class CurrentTriage extends Component {

	intervalID;

	state = {
	  data: "",
	}

	componentDidMount() {
	  this.getData();
	}

	componentWillUnmount() {
	  /*
		stop getData() from continuing to run even
		after unmounting this component. Notice we are calling
		'clearTimeout()` here rather than `clearInterval()` as
		in the previous example.
	  */
	  clearTimeout(this.intervalID);
	}

	getData = () => {
	  fetch(`${Urls.api}/triage`)
		.then(response => response.json())
		.then(data => {
		  this.setState({ data: data});
		  // call getData() again in 5 seconds
		  this.intervalID = setTimeout(this.getData.bind(this), 30000);
		});
	}

	handleClick = channelID => {
		const requestOptions = {
		  method: 'DELETE'
		};
	  
		fetch(`${Urls.api}/triage/`  +  channelID, requestOptions).then((response) => {
		  return response.json();
		}).then((result) => {
		  this.getData()
		});
	  }

	render(){
	
		
		return (

			<div>
			<Table>
				<thead>
				<tr>
				<th>Channel</th>
				<th>Username</th>
				<th>Started</th>
				<th>Delete</th>
				</tr>
			</thead>
			<tbody>
			{this.state.data &&
			 this.state.data.map((item, index) => (
				<tr>
				<th scope="row">{item.Channel.Name}</th>
				<td>{item.User.Name}</td>
				<td>{item.Started}</td>
				<td>
				<button onClick={() => { this.handleClick(item.Channel.ID) }} className="delete-btn">Delete</button>
				</td>
			</tr>
			))}
			
			</tbody>
		  </Table>
			</div>
		)
	}
}

export {CurrentTriage };