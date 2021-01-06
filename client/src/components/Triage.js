import React, { Component  } from 'react';
import axios from 'axios';
import Urls from '../Util/Urls.js';

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
                    this.setState({triage: data.name})
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
        return (
            <div>
                <em>{ triage }</em>
            </div>
        )
    }
}
export {CurrentTriage};