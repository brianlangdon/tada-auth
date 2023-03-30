import {Link, useNavigate} from 'react-router-dom';
import {AUTH_TOKEN} from '../constants';

const Dating = () => {

const authToken = localStorage.getItem(AUTH_TOKEN);

  return (
    <>
    {authToken && (
      <div>
        Dating

        <li>list all the prospective dates</li>
        <li>logged in account is filtered out</li>
        <li>order dates, liked first, no perference second, blocked last</li>

        <Link
              to="/details"
              className="ml1 pointer black"
            >
              Select Details to see details of a possible date
            </Link>

        </div>
    )}
    </>
  )
}

export default Dating;