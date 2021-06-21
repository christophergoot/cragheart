import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export type User = {
  id: number;
  attributes: Record<string, string>;
  events: Record<string, number>;
  last_updated: number;
};

const defaultUser: User = {
  id: 0,
  attributes: {},
  events: {},
  last_updated: 0,
};

const flexStyle = {
  display: "flex",
  justifyContent: "space-between",
  maxWidth: "40rem",
};

export default function Customer() {
  const customerID = useParams<{ id?: string }>().id;
  const [user, setUser] = useState<User>(defaultUser);

  useEffect(() => {
    fetch(`http://localhost:1323/customers/${customerID}`)
      .then((data) => data.json())
      .then((res) => setUser(res.customer));
  }, [customerID]);

  return (
    <div>
      <h2>{user.attributes?.email || "email"}</h2>
      <div>
        <p>{`Last Updated: ${new Date(
          user.last_updated
        ).toLocaleDateString()}`}</p>
      </div>
      <h2>Attributes</h2>
      <div>
        <ul>
          {Object.keys(user.attributes || {}).map((key) => (
            <li style={flexStyle}>
              <div>{key}</div>
              <div>{user.attributes[key]}</div>
            </li>
          ))}
        </ul>
      </div>
      <h2>Events</h2>
      <div>
        <ul>
          <li style={flexStyle}>
            <div>Event Name</div>
            <div>Count</div>
          </li>
          {Object.keys(user.events || {}).map((key) => (
            <li style={flexStyle}>
              <div>{key}</div>
              <div>{user.events[key]}</div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
