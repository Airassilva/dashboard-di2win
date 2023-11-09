import { group } from '@angular/animations';
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'dashboard';
}

window.onload = async () => {
  const dataCount = await getCount("doc_type", 0, null, null, "2023-10-07")
  console.log("Count", dataCount)

  const dataUsers = await getUsers()
  console.log("Users", dataUsers)

  const dataExtracts = await getExtracts()
  console.log("Extracts", dataExtracts)
}

export async function getCount(group_by: string, user_id: number, tipo_documento: string | null, data_comeco: string | null, data_final: string | null) {
  const filtro = `${tipo_documento ? ", tipo_documento: \"" + tipo_documento + "\"" : ""} ${data_comeco ? ", data_comeco: \"" + data_comeco + "\"" : ""} ${data_final ? ", data_final: \"" + data_final + "\"" : ""}`

  const result = await fetch("http://localhost:8080/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        count(group_by: "${group_by}", user_id: ${user_id} ${filtro}) {
          name
          value
        }
      }`
    })
  }).then(res => res.json())

  return result
}

export async function getUsers() {
  const result = await fetch("http://localhost:8080/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        user {
          name
        }
      }`
    })
  }).then(res => res.json())

  return result
}

export async function getExtracts() {
  const result = await fetch("http://localhost:8080/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        extract {
          id
          created_at
          pages_processed
          doc_type
          user_id
        }
      }`
    })
  }).then(res => res.json())

  return result
}
