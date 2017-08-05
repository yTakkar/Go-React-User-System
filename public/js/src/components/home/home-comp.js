import React from 'react'
import $ from 'jquery'
import { Helmet } from 'react-helmet'
import { Link, Redirect } from 'react-router-dom'
import { connect } from 'react-redux'
import { FadeIn } from 'animate-components'
import Timeago from 'handy-timeago'

import * as fn from '../../utils/functions'
import * as user_actions from '../../actions/user-actions'

@connect(store => {
    return {
        user: store.user
    }
})

export default class Notes extends React.Component{

    componentDidMount = () => this.props.dispatch(user_actions.getSession())

    render(){
        let 
            { user: { loggedIn, session: { ID, Username, Email, Joined } } } = this.props,
            j = Timeago(Joined)

        return(
            <div>
                { !loggedIn ? <Redirect to="/login" /> : null }
                <Helmet>
                    <title>Home!!</title>
                </Helmet>
                <FadeIn duration="300ms">
                    <h1>My profile:</h1><br/>
                    <h2>ID: {ID}</h2><br />
                    <h2>Username: {Username}</h2><br/>
                    <h2>Email: {Email}</h2><br />
                    <h2>Joined: {j}</h2><br />
                </FadeIn>
            </div>
        )
    }
}