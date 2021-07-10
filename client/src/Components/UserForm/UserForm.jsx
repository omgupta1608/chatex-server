import LoadingSpinner from '../LoadingSpinner/LoadingSpinner';
import OrButtons from './OrButtons';
import './UserForm.scss';

/**
 * @function
 *
 * @param {Object} props
 * @param {string} props.title form title
 * @param {string} [props.description] form description
 * @param {boolean} [props.showOrButtons] show additional login/register option buttons
 * @param {function} props.onSubmit onSubmit event handler for form
 * @param {boolean} props.isLoading show loading spinner in submit button
 * @param {string} props.submitErrMsg Error message is form submission fails
 * @param {React.ReactNode} props.children form input fields
 */
const UserForm = ({
	title,
	description = '',
	showOrButtons = false,
	onSubmit,
	isLoading = false,
	submitErrMsg = '',
	children,
}) => {
	return (
		<div className='user-form-container'>
			<header className='user-form__header'>
				<h1 className='title'>{title}</h1>
				{description && (
					<p className='desc'>
						<em>{description}</em>
					</p>
				)}
				{submitErrMsg && (
					<p className='submit-err'>
						<em>{submitErrMsg}</em>
					</p>
				)}
			</header>

			<form className='user-form' onSubmit={onSubmit}>
				{children}
				<button type='submit' disabled={isLoading}>
					{title}
					{isLoading && (
						<span className='load-spinner-container'>
							<LoadingSpinner radius='1rem' strokeWidth='0.2rem' />
						</span>
					)}
				</button>
			</form>

			{showOrButtons && <OrButtons prefix={title} />}
		</div>
	);
};

export default UserForm;
