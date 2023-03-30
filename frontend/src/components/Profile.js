import {AUTH_TOKEN} from '../constants';

const Profile = () => {

const authToken = localStorage.getItem(AUTH_TOKEN);

  return (
    <>
    {authToken && (
      <div>
        Profile

        <li>list and edit the profile of the user</li>
        </div>
    )}
    </>
  )
}

export default Profile;