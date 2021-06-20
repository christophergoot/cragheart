import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { User } from "./Customer";

export default function List() {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    fetch(`http://localhost:1323/customers`)
      .then((data) => data.json())
      .then((res) => setUsers(res.customers));
  }, []);

  return (
    <div>
      <h2>Customers</h2>
      <ul>
        {users.map((user) => (
          <li style={{ display: "flex", justifyContent: "space-between" }}>
            <div>{user.id}</div>
            <div>
              <Link to={`/customer/${user.id}`}>
                {user.attributes?.email || "email address"}
              </Link>
            </div>
            <div>{new Date(user.last_updated).toLocaleDateString()}</div>
          </li>
        ))}
      </ul>
    </div>
  );
}
