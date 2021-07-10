import UserForm from '../../Components/UserForm/UserForm';
import UserFormInput from '../../Components/UserForm/UserFormInput';
import UserFormPage from '../../Components/UserFormPage/UserFormPage';

/**
 * user's email verification page
 */
const UserVerificationPage = () => {
	const onSubmit = e => {
		e.preventDefault();
		console.log(e.target.checkValidity());
	};

	return (
		<UserFormPage>
			<UserForm
				title='Verify your Email'
				description={`We sent a 6 digit verification code to your email - email placeholder`}
				onSubmit={onSubmit}
			>
				<UserFormInput
					name='verification_code'
					id='user-verification-code'
					label='Verification Code'
					// TODO: get email from router state
					type='text'
					minLength={6}
					maxLength={6}
					required={true}
				/>
			</UserForm>
		</UserFormPage>
	);
};

export default UserVerificationPage;
