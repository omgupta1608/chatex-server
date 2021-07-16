import UserAuthForm from '../Components/UserAuthForm';
import UserAuthFormInput from '../Components/UserAuthFormInput';
import UserAuthFormPage from '../Components/UserAuthFormPage';
import { Redirect } from 'react-router-dom';
import useAuth from '../Hooks/useAuth';

/**
 * user registeration page
 */
const RegisterPage = () => {
	const {
		state: { errorMsg, isLoading },
		register,
		isUserRegistered,
	} = useAuth();

	const onSubmit = async e => {
		e.preventDefault();
		const formElement = e.target;
		if (!formElement.checkValidity()) return;

		const formData = Object.fromEntries(new FormData(formElement));
		register(formData);
	};

	// redirect to verification page if user already registered
	if (isUserRegistered()) return <Redirect to='/register/verify' />;
	return (
		<UserAuthFormPage>
			<UserAuthForm
				title='Register'
				showOrButtons={true}
				onSubmit={onSubmit}
				isLoading={isLoading}
				submitErrMsg={errorMsg}
			>
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
					type='password'
					minLength={8}
					maxLength={40}
					required={true}
				/>
			</UserAuthForm>
		</UserAuthFormPage>
	);
};

export default RegisterPage;
