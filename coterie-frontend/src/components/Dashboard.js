import React, { useState, useEffect } from 'react';
import SideNav from './SideNav';


const Dashboard = (props) => {
  return (
    <div className="d-flex">
      <SideNav 
      orgID={"1"}
      userID={"1"}
      />
    </div>
  )
}

export default Dashboard;