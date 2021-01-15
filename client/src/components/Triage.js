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

	render(){
	
		
		return (

			<div>
			<Table>
				<thead>
				<tr>
				<th>Channel</th>
				<th>Username</th>
				<th>Started</th>
				</tr>
			</thead>
			<tbody>
			{this.state.data &&
			 this.state.data.map((item, index) => (
				<tr>
				<th scope="row">{item.Channel.Name}</th>
				<td>{item.User.Name}</td>
				<td>{item.Started}</td>
			</tr>
			))}
			
			</tbody>
		  </Table>
			</div>
		)
	}
}

export {CurrentTriage };