import Image from 'next/image'
import Link from "next/link";

async function getData() {
  const res = await fetch('http://localhost:1323/')


  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return res.json()
}

export default async function Page() {
  const data = await getData()

  return (
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        {data.data}
        <Link href={"/time"}>Time</Link>
      </main>
  )
}
