import React, { Component  } from 'react';
import axios from 'axios';
import Urls from '../Util/Urls.js';
import { Table } from 'reactstrap';

class CurrentTriage extends Component {
	constructor(props) {
		super(props);
		this.state = {
			triage: "",
			roles: [],
			error: null,
		}
	}
	componentDidMount() {
		if (this.state.triage === "") {
			this.getCurrentTriage().then(
				data => {
					this.setState({triage: data})
					 this.state.triage.forEach((item, i) => this.state.roles.push(<tr>
						<th scope="row">{item.Channel.name}</th>
						<td>{item.User.name}</td>
						<td>{item.Started}</td>
					</tr>));

				})
			.catch(err => {})
		}
	}

	async getCurrentTriage() {
		const res = await axios.get(`${Urls.api}/triage`);
		console.log(res.data);
		return await res.data;
	}
	
	render(){
		const { roles } = this.state
		
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
				{roles}
			</tbody>
		  </Table>
			</div>
		)
	}
}

export {CurrentTriage };