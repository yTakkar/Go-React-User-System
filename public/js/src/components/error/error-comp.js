import React from 'react'
import $ from 'jquery'
import { Link, Redirect } from 'react-router-dom'
import { Helmet } from 'react-helmet'
import { FadeIn } from 'animate-components'
import { connect } from 'react-redux'
import * as user_actions from '../../actions/user-actions'

@connect(store => {
    return {
        user: store.user
    }
})

export default class Error extends React.Component {

    componentDidMount = () => this.props.dispatch(user_actions.getSession())

    render() {
        let 
            { match: { params: { what } } } = this.props,
            title, desc
            
        if (what == "notfound") {
            title = "User not found"
            desc = "user"
        } else if (what == "note_notfound") {
            title = "Note not found"
            desc = "note"
        } else {
            title = "Error"
            desc = "page"
        }

        return (
            <div class='error' >
                <Helmet>
                    <title>Oops! {title} • Notes App</title>
                </Helmet>
                <FadeIn duration="300ms" >
                    <div className="welcome_div error_div">
                        <div className="error_info">
                            <span>Oops, the {desc} you're looking for does not exist!!</span>
                        </div>
                        <img src="/images/error-3.svg" alt="" />
                        <div class="error_bottom">
                            <Link to='/' className="pri_btn error_login" >Try going to homepage</Link>
                        </div>
                    </div>
                </FadeIn>
            </div>
        )
    }
}
