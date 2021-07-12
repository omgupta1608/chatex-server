import UserAuthForm from '../Components/UserAuthForm';
import UserAuthFormInput from '../Components/UserAuthFormInput';
import UserAuthFormPage from '../Components/UserAuthFormPage';

/**
 * user registeration page
 */
const RegisterPage = () => {
	const onSubmit = e => {
		e.preventDefault();
		console.log(e.target.checkValidity());
	};

	return (
		<UserAuthFormPage>
			<UserAuthForm title='Register' showOrButtons={true} onSubmit={onSubmit}>
				<UserAuthFormInput
					name='name'
					id='register-name'
					label='Name'
					type='text'
					minLength={3}
					maxLength={15}
					required={true}
					autoFocus={true}
				/>
				<UserAuthFormInput
					name='email'
					id='register-email'
					label='E-Mail'
					type='email'
					required={true}
				/>
				<UserAuthFormInput
					name='password'
					id='register-password'
					label='Password'
					autoComplete='new-password'
					minLength={8}
					maxLength={40}
					required={true}
				/>
			</UserAuthForm>
		</UserAuthFormPage>
	);
};

export default RegisterPage;
