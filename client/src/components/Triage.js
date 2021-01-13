import React, { Component  } from 'react';
import axios from 'axios';
import Urls from '../Util/Urls.js';
import { Table } from 'reactstrap';

class CurrentTriage extends Component {
	constructor(props) {
		super(props);
		this.state = {
			triage: "",
			error: null,
		}
	}
	componentDidMount() {
		if (this.state.triage === "") {
			this.getCurrentTriage().then(
				data => {
					this.setState({triage: data})
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
		const { triage } = this.state
		const items = this.state.triage.map((item) =>
			<tr>
				<th scope="row">{item.Channel.name}</th>
				<td>{item.User.name}</td>
				<td>{item.Started}</td>
			</tr>
		);
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
				{items}
			</tbody>
		  </Table>
			</div>
		)
	}
}

export {CurrentTriage };