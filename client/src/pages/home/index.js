import React from 'react'
import Filter from '../../components/filter'
import ItemList from '../../components/itemList'

export default function Home() {
  return (
    <div>
      <div>
        <Filter />
      </div>
      <div>
        <ItemList />
      </div>
    </div>
  )
}
