import UserAuthForm from '../Components/UserAuthForm';
import UserAuthFormInput from '../Components/UserAuthFormInput';
import UserAuthFormPage from '../Components/UserAuthFormPage';

/**
 * user login page
 */
const LoginPage = () => {
	const onSubmit = e => {
		e.preventDefault();
		console.log(e.target.checkValidity());
	};

	return (
		<UserAuthFormPage>
			<UserAuthForm
				title='Login'
				showOrButtons={true}
				onSubmit={onSubmit}
				description='Login with your email and password'
			>
				<UserAuthFormInput
					name='email'
					id='login-email'
					label='E-Mail'
					type='email'
					required={true}
					autofocus={true}
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
