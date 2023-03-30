import {AUTH_TOKEN} from '../constants';

const Details = () => {

const authToken = localStorage.getItem(AUTH_TOKEN);

  return (
    <>
    {authToken && (
      <div>
        Details

        <li>list the details of the selected date</li>
        </div>
    )}
    </>
  )
}

export default Details;