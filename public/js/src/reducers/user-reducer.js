const user_def = {
    loggedIn: false,
    session: {
        ID: null,
        Username: null,
        Email: null,
        Joined: null
    }
}

const user = (state=user_def, action) => {
    let py = action.payload
    switch (action.type) {
        case "LOGIN":
            return { ...state, loggedIn: py.loggedIn, session: py.session }
            break

        case "GET_SESSION":
            return { ...state, loggedIn: py.loggedIn, session: py.session }
            break

        case "LOGOUT":
            return { ...state, loggedIn: false, session: {} }
            break
    }
    return state
}

export default user