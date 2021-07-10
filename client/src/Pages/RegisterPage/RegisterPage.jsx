import UserForm from '../../Components/UserForm/UserForm';
import UserFormInput from '../../Components/UserForm/UserFormInput';
import UserFormPage from '../../Components/UserFormPage/UserFormPage';

/**
 * user registeration page
 */
const RegisterPage = () => {
	const onSubmit = e => {
		e.preventDefault();
		console.log(e.target.checkValidity());
	};

	return (
		<UserFormPage>
			<UserForm title='Register' showOrButtons={true} onSubmit={onSubmit}>
				<UserFormInput
					name='name'
					id='register-name'
					label='Name'
					type='text'
					minLength={3}
					maxLength={15}
					required={true}
				/>
				<UserFormInput
					name='email'
					id='register-email'
					label='E-Mail'
					type='email'
					required={true}
				/>
				<UserFormInput
					name='password'
					id='register-password'
					label='Password'
					autoComplete='new-password'
					minLength={8}
					maxLength={40}
					required={true}
				/>
			</UserForm>
		</UserFormPage>
	);
};

export default RegisterPage;
