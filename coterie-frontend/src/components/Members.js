import React, {useState, useEffect} from 'react';
import Member from './Member';
import axios from 'axios';

const Members = (props) => {
  const [membersList, setMembersList] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/members`

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setMembersList(list);
      })
      .catch((error) => {
        console.log(`There was an error retrieving members: ${error}`)
      });
  },[url])

  let memberComponents = undefined
  if (membersList !== null && membersList.length > 0) {
    memberComponents = membersList.map((mem) => {
    return(
      <Member
      key={mem.id}
      id={mem.id}
      name={mem.name}
      birthdate={mem.birthdate}
      class={mem.class}
      email={mem.email}
      fundsRaised={mem.funds_raised}
      orgID={mem.organization_id}
      />
    )
  })
} else {
  memberComponents = <p className="open-sans">It looks like you don't have any members, yet!</p>
}


  return (
    <div className="">
      <h6 className="text-left w-100">Followers</h6>
      <table className="table table-hover table-light">
        <thead className="thead-light text-center">
          <tr>
            <th>Name</th>
            <th>Class</th>
            <th>Raised</th>
            <th>Email</th>
            <th>Birthdate</th>
          </tr>
        </thead>
        <tbody>
          {memberComponents}
        </tbody>
      </table>
    </div>
  )
}

export default Members;

