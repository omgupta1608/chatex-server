import { useEffect, useReducer } from 'react';

/**
 * @typedef {Object} USER
 * @property {string} [uid]
 * @property {string} [name]
 * @property {string} [about]
 * @property {string} [email]
 * @property {string} [password]
 * @property {string} [profile_pic]
 */

export const ACTION_TYPES = {
	INIT: 'INIT',
	ERROR: 'ERROR',
	SUCCESS: 'SUCCESS',
	RESET: 'RESET',
};

/**
 * @typedef {Object} STATE
 * @property {boolean} isLoading
 * @property {string} errorMsg
 * @property {?USER} user
 */
const initState = {
	isLoading: false,
	errorMsg: '',
	// TODO: get user data from db
	user: null,
};

/**
 *
 * @param {STATE} _
 * @param {{type: string, payload: any}} action
 * @returns {STATE}
 */
const reducer = (_, action) => {
	switch (action.type) {
		case ACTION_TYPES.INIT:
			return { isLoading: true, errorMsg: '', user: null };
		case ACTION_TYPES.ERROR:
			return { isLoading: false, errorMsg: action.payload, user: null };
		case ACTION_TYPES.SUCCESS:
			return { isLoading: false, errorMsg: '', user: action.payload };
		case ACTION_TYPES.RESET:
			return initState;
		default:
			throw new Error('Invalid action type');
	}
};

const serverURL = process.env.REACT_APP_API_URL;
const abortController = new AbortController();
const useProvideAuth = () => {
	const [state, dispatch] = useReducer(reducer, initState);

	// abort fetch on unmount
	useEffect(() => () => abortController.abort(), []);

	/**
	 *
	 * @param {string} url
	 * @param {any} reqBody
	 * @returns {boolean} if request returned without error
	 */
	const postData = async (url, reqBody) => {
		dispatch({ type: ACTION_TYPES.INIT });

		let resData;
		try {
			const rawRes = await fetch(url, {
				mode: 'cors',
				method: 'post',
				headers: {
					'Content-Type': 'application/json',
					Accept: 'application/json',
				},
				body: JSON.stringify(reqBody),
				signal: abortController.signal,
			});

			resData = await rawRes.json();
		} catch (err) {
			if (err.name !== 'AbortError')
				dispatch({
					type: ACTION_TYPES.ERROR,
					payload: 'Something went wrong. Try again in a few minutes',
				});
			return false;
		}

		const { error, message, user, jwt } = resData;
		if (error) {
			dispatch({
				type: ACTION_TYPES.ERROR,
				payload: `${message} - ${error}`,
			});
			return false;
		}

		if (jwt) user.jwt = jwt;
		dispatch({
			type: ACTION_TYPES.SUCCESS,
			payload: user,
		});
		return true;
	};

	/**
	 * login user
	 * @param {Object} reqData
	 * @param {string} reqData.email user's email id
	 * @param {string} reqData.password user's password
	 */
	const login = async reqData => {
		const loginURL = new URL('./login', serverURL).href;
		return postData(loginURL, reqData);
		// TODO: save user to db if success
	};

	/**
	 * register user
	 * @param {Object} reqData
	 * @param {string} reqData.name user's name
	 * @param {string} reqData.email user's email id
	 * @param {string} reqData.password user's password
	 */
	const register = async reqData => {
		const registerURL = new URL('./register', serverURL).href;
		return postData(registerURL, reqData);
	};

	/**
	 * verify user using uid and verification code
	 * @param {Object} reqData
	 * @param {string} reqData.uid user's uid
	 * @param {string} reqData.verificationCode user's verification code
	 */
	const verifyUserAndLogin = async reqData => {
		const verifyUser = new URL('./register/verify', serverURL).href;
		return postData(verifyUser, reqData);
		// TODO: save user to db if success
	};

	const logout = async () => {
		// TODO: clear db
		dispatch({ type: ACTION_TYPES.RESET });
		return true;
	};

	/**
	 * @returns {boolean} if user is authenticated
	 */
	const isUserAuthenticated = () => {
		const { user } = state;
		return user && user.jwt;
	};

	/**
	 * @returns {boolean} if user is registered and waiting for user verification
	 */
	const isUserRegistered = () => {
		const { user } = state;
		return user && user.uid;
	};

	return {
		state,
		login,
		register,
		verifyUserAndLogin,
		logout,
		isUserAuthenticated,
		isUserRegistered,
	};
};

export default useProvideAuth;
