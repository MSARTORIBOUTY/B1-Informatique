const blockChain = (data, prev = { index: 0, hash: '0' }) => {
    let obj  = {
        data,
        prev,
    }

    obj.index = obj.prev.index + 1
    // La méthode JSON.stringify() convertit une valeur JavaScript en chaîne JSON.
    // concaténation d'index, hash and data en stringify
    obj.hash = hashCode(obj.index + prev.hash + JSON.stringify(data))
    // chain a function that takes a new data and create the next block with it.
    obj.chain = (data) => blockChain(data,obj)

  return obj
}