const App = () => {
  return (
    <section className="p-4">
      <h2 className="text-3xl font-bold mb-4">Not Just Potato Phone <div className="text-md text-blue-600">Project</div></h2>
      <ProjectCard
        description="A private (possibly public later) project aimed at turning the Samsung J6PRIMELTE into an SSH server for experimentation and pushing beyond limitations. 🚀"
      />
      <div className="mt-4">
        <h3 className="text-xl font-bold">Maintainers</h3>
        <div className="flex space-x-4">
          <MaintainerCard name="mac" id="829156179803504670" displayname="bluestar.png"/>
          <MaintainerCard name="fuse" id="736163902835916880" displayname="blueskychan_"/>
        </div>
      </div>
      <div className="mt-4">
        <h3 className="text-xl font-bold">Status</h3>
        <p className="mb-2">Status Page(Near Realtime): <a href="https://status.fusemeow.codes" className="text-blue-500 underline">Check here</a></p>
        <p>Data API (REST API): <a href="https://api.fusemeow.codes/data" className="text-blue-500 underline">Access here</a></p>
      </div>
    </section>
  );
};

const ProjectCard = ({ description }) => {
  return (
    <div className="border-2 border-zinc-400 p-4 rounded-lg shadow-md w-fit">
      <p className="text-sm text-gray-300">{description}</p>
    </div>
  );
};

const MaintainerCard = ({ name, id, displayname }) => {
  const discordUrl = `https://discord.com/users/${id}`;

  return (
    <div className="border-2 border-zinc-400 p-2 rounded-lg shadow-sm">
      <div className="text-lg font-semibold">{name}</div>
      <div className="text-sm text-gray-300">
        Discord: <a href={discordUrl} className="text-blue-500 underline">{displayname}</a>
      </div>
    </div>
  );
};

export default App;

