import React, { useState, useEffect } from 'react';
import SideNav from './SideNav';
import Organization from './Organization';
import axios from 'axios';


const Dashboard = () => {
  //TODO: logic that checks the URL... is it just the dashboard? then show xyz stuff.
  const [organization, setOrganization] = useState(null);

  const url = "http://localhost.com/3000/organization" 
  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const organization = response.headers;
        setOrganization(organization);
      })
      .catch((error) => {
        console.log(`There was an error: ${error}`)
      });
  },[url])


  let orgComponent = undefined
  if (organization !== null) {
    orgComponent = <Organization
      key = {organization.id}
      id = {organization.id}
      name = {organization.name}
      mission_statement = {organization.mission_statement}
      total_funds = {organization.total_funds}
      user_id = {organization.user_id}
    /> } else {
    orgComponent = <p className="open-sans">It looks like you haven't added your org yet!</p>
  }



  return (
    <div className="d-flex">
      
      <div className="">
        <SideNav />
        {orgComponent}
        
      </div>
    </div>
  )
}

export default Dashboard;