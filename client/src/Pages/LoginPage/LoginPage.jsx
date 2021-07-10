import UserForm from '../../Components/UserForm/UserForm';
import UserFormInput from '../../Components/UserForm/UserFormInput';
import UserFormPage from '../../Components/UserFormPage/UserFormPage';

/**
 * user login page
 */
const LoginPage = () => {
	const onSubmit = e => {
		e.preventDefault();
		console.log(e.target.checkValidity());
	};

	return (
		<UserFormPage>
			<UserForm
				title='Login'
				showOrButtons={true}
				onSubmit={onSubmit}
				description='Login with your email and password'
			>
				<UserFormInput
					name='email'
					id='login-email'
					label='E-Mail'
					type='email'
					required={true}
				/>
				<UserFormInput
					name='password'
					id='login-password'
					label='Password'
					autoComplete='current-password'
					minLength={8}
					maxLength={40}
					required={true}
				/>
			</UserForm>
		</UserFormPage>
	);
};

export default LoginPage;
