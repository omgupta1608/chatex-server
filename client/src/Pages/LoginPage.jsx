import UserAuthForm from '../Components/UserAuthForm';
import UserAuthFormInput from '../Components/UserAuthFormInput';
import UserAuthFormPage from '../Components/UserAuthFormPage';
import useAuth from '../Hooks/useAuth';

/**
 * user login page
 */
const LoginPage = () => {
	// default redirection is to dashboard
	const {
		state: { errorMsg, isLoading },
		login,
	} = useAuth();

	const onSubmit = async e => {
		e.preventDefault();
		const formElement = e.target;
		if (!formElement.checkValidity()) return;

		const formData = Object.fromEntries(new FormData(formElement));

		login(formData);
	};

	return (
		<UserAuthFormPage>
			<UserAuthForm
				title='Login'
				showOrButtons={true}
				onSubmit={onSubmit}
				submitErrMsg={errorMsg}
				isLoading={isLoading}
				description='Login with your email and password'
			>
				<UserAuthFormInput
					name='email'
					id='login-email'
					label='E-Mail'
					type='email'
					required={true}
					autoFocus={true}
				/>
				<UserAuthFormInput
					name='password'
					id='login-password'
					label='Password'
					type='password'
					autoComplete='current-password'
					minLength={8}
					maxLength={40}
					required={true}
				/>
			</UserAuthForm>
		</UserAuthFormPage>
	);
};

export default LoginPage;
