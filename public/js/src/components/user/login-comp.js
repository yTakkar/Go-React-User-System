import React from 'react'
import $ from 'jquery'
import { Helmet } from 'react-helmet'
import { Link, Redirect } from 'react-router-dom'
import { connect } from 'react-redux'
import { FadeIn } from 'animate-components'

import * as fn from '../../utils/functions'
import * as user_actions from '../../actions/user-actions'

@connect(store => {
    return {
        user: store.user
    }
})

export default class Login extends React.Component {

    state = {
        username: "",
        password: "",
        fnLoggedIn: false
    }

    componentDidMount = () => this.props.dispatch(user_actions.getSession())

    update_ = (e, of) => {
        let v = e.target.value
        switch (of) {
            case "username":
                this.setState({ username: v })
                break;
            case "password":
                this.setState({ password: v })
                break;
        }
    }

    login = e => {
        e.preventDefault()
        let
            username = $('.l_username').val(),
            password = $('.l_password').val()

        if (!username || !password) {
            Notify({ value: "Values are missing!" })
        } else {

            let loginOpt = {
                of: "login",
                dispatch: this.props.dispatch,
                data: {
                    username: $('.l_username').val(),
                    password: $('.l_password').val()
                },
                btn: $('.l_submit'),
                url: "/user/login",
                done: () => this.setState({ fnLoggedIn: true }),
                defBtnValue: "Login to continue",
            }
            fn.commonLogin(loginOpt)

        }
    }

    render() {
        let 
            { username, password, fnLoggedIn } = this.state,
            { user: { loggedIn } } = this.props

        return (
            <div>
                { (fnLoggedIn || loggedIn) ? <Redirect to="/" /> : null }
                <Helmet>
                    <title>Login To Continue!!</title>
                </Helmet>
                <FadeIn duration="300ms">
                    <div class="notes_wrapper">
                        <div class="log_sign">
                            <Link to="/signup" class="pri_btn">Need an account?</Link>
                        </div>
                        <div class="register cua" >
                            <div class="display_text">
                                <span>Get started again</span>
                            </div>
                            <form class="form_login" onSubmit={this.login} >
                                <input 
                                    type="text" 
                                    name="username" 
                                    value={username} 
                                    class="l_username" 
                                    autoFocus 
                                    required 
                                    spellCheck="false" 
                                    autoComplete='false' 
                                    placeholder='Username' 
                                    onChange={e => this.update_(e, "username")}
                                />
                                <input 
                                    type="password" 
                                    name="password" 
                                    value={password} 
                                    class="l_password" 
                                    required 
                                    placeholder='Password' 
                                    onChange={e => this.update_(e, "password")}
                                />
                                <input type="submit" name="" value="Login to continue" class="l_submit" />
                            </form>
                        </div>
                    </div>
                </FadeIn>
            </div>
        )
    }
}
