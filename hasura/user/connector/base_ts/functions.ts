import { Client } from "pg";
import dotenv from "dotenv";

// Load environment variables from .env.local
dotenv.config({ path: ".env.local" });

/**
 * @readonly Exposes the function as an NDC function (the function should only query data without making modifications)
 */
export function hello(name?: string) {
  return `hello ${name ?? "world"}`;
}

/**
 * @param userData An object containing the user's data.
 * @returns The ID of the newly inserted user.
 */
export async function insertUser(userData: { username: string; password: string }): Promise<string> {
  const client = new Client({
    connectionString: process.env.CONNECTION_URI,
  });

  console.log(process.env.CONNECTION_URI);

  await client.connect();

  const queryText = `
    INSERT INTO users (username, password)
    VALUES ($1, $2)
    RETURNING id
  `;
  const values = [userData.username, userData.password];
  const result = await client.query(queryText, values);

  await client.end();

  if (result.rows.length > 0) {
    return result.rows[0].id;
  } else {
    throw new Error("Failed to insert user");
  }
}
